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
      console.log(rooms)
      $.ajax({
        url: '/getDates',
        type: "get",
        dataType: 'json',
      }).then(function (dates) {
        console.log(rooms)

        $('#thead').append(
          '<tr id="tr_dates">'
          + '<th>No</th>'
          + '<th>Room</th>'
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
            + '<th>1'
            + '</th>'
            + '<th>'
            + '<a href="" title="Leicester City F.C.">' + room_name + '</a>'
            // + '<td>' + room_name + '</td>'
            + '</th>'
            + '</tr>'

          );
          $.ajax({
            url: '/getDates',
            type: "get",
            dataType: 'json',
          }).then(function (dates) {
            for (let key in dates) {
              $('#' + id).append(
                '<td>' + dates[key] + '</td>'
              )
            }
          });
        });
      }
    });
  }
}
const tableClass = new table();

tableClass.createTable()