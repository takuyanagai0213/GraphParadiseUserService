package interactor

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/takuyanagai0213/GraphParadiseUserService/authorization"
	// "github.com/takuyanagai0213/GraphParadiseUserService/grpc/postservice"

	"github.com/go-playground/validator/v10"
	"github.com/takuyanagai0213/GraphParadiseUserService/db"
	"github.com/takuyanagai0213/GraphParadiseUserService/domain/model"
	"github.com/takuyanagai0213/GraphParadiseUserService/usecase/repository"
	"golang.org/x/crypto/bcrypt"
)

type key int

var (
	err  error
	user model.User
	// relation  model.Relation
	users []model.User
	// relations []model.Relation
	rows     *sql.Rows
	validate *validator.Validate
)

const (
	// secretを環境変数から読む
	secret = "2FMd5FNSqS/nW2wWJy5S3ppjSHhUnLt8HuwBkTD6HqfPfBBDlykwLA=="
	// キータイプ
	stringKey key = iota
	// ゼロ値
	zero uint32 = 0
	// one 1
	one uint32 = 1
)

const (
	// authorityNormalUser 一般ユーザー
	authorityNormalUser uint32 = 1
	// authoritySuperUser 管理者ユーザー
	authoritySuperUser uint32 = 9
)

var demoUser = model.User{UserName: "", Email: "", Password: "password", Authority: authorityNormalUser}
var demoSuperUser = model.User{UserName: "", Email: "", Password: "password", Authority: authoritySuperUser}

// UserInteractor ユーザサービスを提供するメソッド群
type UserInteractor struct{}

var _ repository.UserRepository = (*UserInteractor)(nil)

// Create ユーザ1件を作成
func (i *UserInteractor) Create(postData *model.User) (*model.User, error) {
	validate = validator.New()
	// DB := db.GetDB()
	createUser := postData

	u, _ := i.GetUserByEmail(createUser.Email)
	if u.ID != 0 {
		return postData, errors.New("email already used")
	}

	// User構造体のバリデーション
	if err := validate.Struct(postData); err != nil {
		return postData, err
	}
	inputPassword := postData.Password

	hash, err := createHashPassword(inputPassword)
	createUser.Password = hash

	if err != nil {
		return createUser, err
	}

	// トランザクション開始
	tx := db.StartBegin()

	if err := tx.Create(createUser).Error; err != nil {
		db.EndRollback()
		return postData, err
	}

	// トランザクションを終了しコミット
	db.EndCommit()
	return postData, nil
}

// DeleteByID 指定したIDのユーザー1件を削除
func (i *UserInteractor) DeleteByID(id uint32) error {
	var user model.User

	// トランザクション開始
	tx := db.StartBegin()

	if err := tx.Where("id = ? ", id).Delete(&user).Error; err != nil {
		db.EndRollback()
		return err
	}
	// if err = deletePostsCommentsByUserID(id); err != nil {
	// 	db.EndRollback()
	// 	return err
	// }

	// トランザクションを終了しコミット
	db.EndCommit()
	return nil
}

// Count ユーザ件数を取得
func (i *UserInteractor) Count(user model.User) (int, error) {
	var count int
	DB := db.GetDB()
	if err := DB.Find(&user).Count(&count).Error; err != nil {
		return count, err
	}
	return count, nil
}

// List ユーザを全件取得
func (i *UserInteractor) List() ([]model.User, error) {
	var userList []model.User
	rows, err := listAll(context.Background())
	if err != nil {
		fmt.Println("Error happened")
		return []model.User{}, err
	}
	for _, row := range rows {
		userList = append(userList, row)
	}

	return userList, nil
}

// ListAllNormalUser 一般ユーザーリストを返却
func (i *UserInteractor) ListAllNormalUser() ([]model.User, error) {
	DB := db.GetDB()
	var users []model.User
	err := DB.Where("authority = ?", authorityNormalUser).Select("users.id, users.user_name, users.profile_text, users.authority").Find(&users).Error
	if err != nil {
		fmt.Println("Error happened")
		return []model.User{}, err
	}

	return users, nil
}

// listAll 全件取得
func listAll(ctx context.Context) ([]model.User, error) {
	DB := db.GetDB()

	rows, err := DB.Find(&users).Rows()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		DB.ScanRows(rows, &user)
		users = append(users, user)
	}
	return users, nil
}

// Update ユーザを更新する
func (i *UserInteractor) Update(postData *model.User) (*model.User, error) {
	validate = validator.New()
	DB := db.GetDB()
	// postされたIdに紐づくuserを取得
	id := postData.ID
	findUser := &model.User{}

	if err := DB.Where("id = ?", id).First(&findUser).Error; err != nil {
		log.Fatalf("err: %v", err)
		return findUser, err
	}

	// User構造体のバリデーション
	if err := validate.Struct(postData); err != nil {
		return postData, err
	}

	updateUser := postData

	// 更新時に入力されたemailが他のユーザーと重複していないか判定
	if i.OtherUserExistsByEmail(updateUser.Email, updateUser.ID) == true {
		return postData, errors.New("email already used")
	}
	if postData.Password != "" {
		// // パスワードをhash
		hash, err := createHashPassword(postData.Password)
		// hashしたパスワードをSaveするuserにセット
		updateUser.Password = string(hash)
		if err != nil {
			return updateUser, err
		}
	}

	updateUser.ID = findUser.ID

	// トランザクション開始
	tx := db.StartBegin()

	if err := tx.Model(&user).Update(&postData).Error; err != nil {
		db.EndRollback()
		return updateUser, err
	}
	// トランザクションを終了しコミット
	db.EndCommit()
	return updateUser, nil
}

// GetUserByUserID UserIDを元にユーザを1件取得する
func (i *UserInteractor) GetUserByUserID(id uint32) (model.User, error) {
	var user model.User

	DB := db.GetDB()
	row := DB.Where("id = ?", id).First(&user)
	if err := row.Error; err != nil {
		return user, err
	}
	DB.Table(db.UserTableName).Scan(row)

	return user, nil
}

// GetUserByEmail Emailを元にユーザを1件取得する
func (i *UserInteractor) GetUserByEmail(email string) (model.User, error) {
	var user model.User

	DB := db.GetDB()
	row := DB.Where("email = ?", email).First(&user)
	if err := row.Error; err != nil {
		return user, err
	}
	DB.Table(db.UserTableName).Scan(row)

	return user, nil
}

// OtherUserExistsByEmail 指定したユーザ以外にemailが重複するユーザが存在するか判定
func (i *UserInteractor) OtherUserExistsByEmail(email string, id uint32) bool {
	var user model.User
	var count int

	DB := db.GetDB()
	DB.Not("id = ?", id).Where("email = ?", email).Find(&user).Count(&count)
	return count > 0
}

func createHashPassword(password string) (string, error) {
	// パスワードの暗号化
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	hashPassword := string(hash)

	if err != nil {
		log.Fatal(err)
		return hashPassword, err
	}
	return hashPassword, nil
}

// LoginAuth パスワード入力による認証メソッド
func (i *UserInteractor) LoginAuth(email string, inputPassword string) (*model.Auth, error) {
	// 入力ユーザ存在有無の判定.
	// eメールに紐づくユーザのパスワードを取得
	findUser, err := i.GetUserByEmail(email)
	if err != nil {
		return &model.Auth{}, err
	}

	// DBから取得したパスワードと入力値のハッシュを比較
	err = bcrypt.CompareHashAndPassword([]byte(findUser.Password), []byte(inputPassword))
	// 認証失敗
	if err != nil {
		return &model.Auth{}, err
	}

	// contextにユーザ情報格納
	ctx := context.Background()
	ctx = context.WithValue(ctx, stringKey, findUser.ID)

	// authのAuthFuncを呼び出す
	// jwt生成
	token, err := authorization.CreateToken(&findUser)
	if err != nil {
		return &model.Auth{}, err
	}

	return &model.Auth{
		UserID:    findUser.ID,
		Authority: findUser.Authority,
		Token:     token,
	}, nil
}

// CreateDemoUser ゲストログインユーザーを作成し、トークンを返す
func (i *UserInteractor) CreateDemoUser() (*model.Auth, error) {
	// ユーザーIDの最大値を取得
	maxID := getMaxUserID()
	maxID++

	// maxIDをdemoユーザーのEmailに格納
	uniqueDemoUser := demoUser

	uniqueDemoUser.UserName = "ゲストユーザー" + strconv.Itoa(maxID)
	uniqueDemoUser.Email = strconv.Itoa(maxID) + "demouser@example.com"
	// Create実行
	_, err := i.Create(&uniqueDemoUser)

	if err != nil {
		return &model.Auth{}, err
	}
	// LoginAuth実行
	auth, err := i.LoginAuth(uniqueDemoUser.Email, "password")
	if err != nil {
		return &model.Auth{}, err
	}
	// 認証情報を返す
	return auth, nil
}

// CreateDemoSuperUser 管理者ユーザーを作成し、トークンを返す
func (i *UserInteractor) CreateDemoSuperUser() (*model.Auth, error) {
	// ユーザーIDの最大値を取得
	maxID := getMaxUserID()
	maxID++

	// maxIDをdemoユーザーのEmailに格納
	uniqueDemoSuperUser := demoSuperUser

	uniqueDemoSuperUser.UserName = "管理者ユーザー" + strconv.Itoa(maxID)
	uniqueDemoSuperUser.Email = strconv.Itoa(maxID) + "superuser@example.com"
	// Create実行
	_, err := i.Create(&uniqueDemoSuperUser)

	if err != nil {
		return &model.Auth{}, err
	}
	// LoginAuth実行
	auth, err := i.LoginAuth(uniqueDemoSuperUser.Email, "password")
	if err != nil {
		return &model.Auth{}, err
	}
	// 認証情報を返す
	return auth, nil
}

func getMaxUserID() int {
	var result int
	DB := db.GetDB()
	row := DB.Table(db.UserTableName).Select("max(id)").Row()

	row.Scan(&result)
	return result
}

// // Follow ユーザーをフォロー
// func (i *UserInteractor) Follow(postData *model.Relation) (*model.Relation, error) {
// 	DB := db.GetDB()
// 	if err := DB.Create(postData).Error; err != nil {
// 		return postData, err
// 	}
// 	return postData, nil
// }

// // UnFollow フォロー解除
// func (i *UserInteractor) UnFollow(postData *model.Relation) (*model.Relation, error) {
// 	DB := db.GetDB()
// 	if err := DB.Where("follower_user_id = ?", postData.FollowerUserID).Where("followed_user_id = ?", postData.FollowedUserID).Delete(postData).Error; err != nil {
// 		return postData, err
// 	}
// 	return postData, nil
// }

// GetFollowUsersByID フォロワーユーザーIDを一覧で取得
// func (i *UserInteractor) GetFollowUsersByID(userID uint32) []uint32 {

// 	var followers []uint32
// 	DB := db.GetDB()
// 	rows, err := DB.Where("followed_user_id = ?", userID).Find(&relations).Rows()
// 	if err != nil {
// 		log.Println("Error occured on GetFollowUsersByID", userID)
// 		return nil
// 	}
// 	for rows.Next() {
// 		DB.ScanRows(rows, &relation)
// 		followers = append(followers, relation.FollowerUserID)
// 	}

// 	return followers
// }

// countFollowUserByFollower フォロワーユーザーIDを元にフォローされている数を取得する
// func countFollowUserByFollower(ID uint32) int {
// 	var count int
// 	DB := db.GetDB()
// 	DB.Where("follower_user_id = ?", ID).Model(&relation).Count(&count)

// 	return count
// }

// deletePostsByUserID 退会したユーザーの投稿・コメントを削除
// func deletePostsCommentsByUserID(userID uint32) error {
// 	postURL := os.Getenv("POST_URL")
// 	cc, err := grpc.Dial(postURL, grpc.WithInsecure())
// 	if err != nil {
// 		log.Fatalf("could not connect: %v", err)
// 	}

// 	defer cc.Close()
// 	postClient := postservice.NewPostServiceClient(cc)

// 	request := &postservice.DeletePostsCommentsByUserIDRequest{
// 		CreateUserId: userID,
// 	}

// 	_, err = postClient.DeletePostsCommentsByUserID(context.Background(), request)
// 	return err
// }
