function User() {
  this.view_s = $.extend(true, {}, this.def_s);
  return;
}
User.prototype.def_s = new Object({
  rooms: {},
  room_name_list: [],
  users: []
});
User.prototype.init = function() {
  this.showUser();
}
User.prototype.showUser = function() {
  $.ajax({
    url: '/user/get',
    type: "get",
    dataType: 'json',
  }).then(this.responseData.bind(this));
}
User.prototype.responseData = function(users) {
  this.view_s.users = users;
  this.createUserList();
}
User.prototype.createUserList = function(message) {
  for( let key in this.view_s.users){
    $('#user_list').append(
      '<ol>'
      + '<li>' + this.view_s.users[key]['ID'] + '</li>'
      + '<li>' + this.view_s.users[key]['Name'] + '</li>'
      + '</ol>'
      );
  }
};
  
const userClass = new User();

userClass.init()