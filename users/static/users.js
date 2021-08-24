function User() {
  this.view_s = $.extend(true, {}, this.def_s);
  return;
}
User.prototype.def_s = new Object({
  rooms: {},
  room_name_list: []
});
User.prototype.init = function() {
  $('#submit').on('click', this.createUser.bind(this))
}
User.prototype.createUser = function() {
  $.ajax({
    url: '/user/new',
    type: "get",
    dataType: 'json',
  }).then(this.responseData.bind(this));
}
User.prototype.responseData = function(message) {
  console.log(message)
}
const userClass = new User();

userClass.init()