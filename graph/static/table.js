var rooms = [];
class table {
  constructor(){
    this.rooms = []
  }
  createTable() {
    const self = this;
    $.ajax({
      url: '/getData1',
      type: "get",
      dataType: 'json',
    }).then(function (json) {
      $.ajax({
        url: '/getRooms',
        type: "get",
        dataType: 'json',
      }).then(function (rooms) {
        self.drowChart(json, rooms, "myChart1");
        self.drowLineChart(json, rooms, "myChart2");
        self.drowPieChart(json, rooms, "myChart3");
        self.drowRadarChart(json, rooms, "myChart4");
      });
    });
    $.ajax({
      url: '/getDataForDaily',
      type: "get",
      dataType: 'json',
    }).then(function (json) {
      $.ajax({
        url: '/getDates',
        type: "get",
        dataType: 'json',
      }).then(function (dates) {
        self.drowLineChart(json, dates, "myChart2");
      });
    });
  }
}
const tableClass = new table();

tableClass.createTable()