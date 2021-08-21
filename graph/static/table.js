var rooms = [];
class table {
  constructor(){
    this.rooms = []
  }
  createTable() {
    const self = this;
    $.ajax({
      url: '/getRooms',
      type: "get",
      dataType: 'json',
    }).then(function (rooms) {
      $.ajax({
        url: '/getDates',
        type: "get",
        dataType: 'json',
      }).then(function (dates) {
        console.log(rooms)

        $('#thead').append(
          '<tr id="tr_dates">'
          + '<th></th>'
          + '<tr>'
        );
        for (let key in dates) {
          $('#tr_dates').append(
            '<td>' + dates[key] + '日</td>'
          )
        }
      });
      for (let key in rooms) {
        const room_name = rooms[key];
        $.ajax({
          url: '/GetDataForTable',
          type: "get",
          data: room_name,
          dataType: 'json',
        }).then(function (json) {
          const id = 'tr_' + room_name;
          $('#tbody').append(
            '<tr id=' + id + '>'
            + '</tr>'

          );

          $('#' + id).append(
            '<th>' + room_name + '</th>'
            +    '<td>2</td>'
            +    '<td>2</td>'
            +    '<td>2</td>'
            +    '<td>2</td>'
            +    '<td>2</td>'
            +    '<td>2</td>'
          )
        });
      }
    });
  }
}
const tableClass = new table();

tableClass.createTable()