function autoRefresh() {
      window.location = window.location.href;  
}

function getrefresh_timeout() {
      let today = new Date();

      let hours= today.getHours();
      let minutes = today.getMinutes();
      let seconds = today.getSeconds();

      /*Refresh time is 1 minute after midnight*/
      let refresh_timeout = (86400 - (hours*3600 + (minutes+1)*60 + seconds)) * 1000;
      return refresh_timeout;
}

function getDailyFileName(){
      let today = new Date();
      console.log(today.toJSON())


      function gfg_Run() {
            var date = today.toJSON().slice(0, 10);
            console.log(date)
            var nDate = date.slice(8, 10) + '/' 
                       + date.slice(5, 7) + '/' 
                       + date.slice(0, 4);
            console.log(nDate);
      }
      gfg_Run()

      let year = today.getFullYear();
      let month = today.getMonth()+1;
      if (month<10)mm = '0'+month;
      let day = today.getDate();
      if (day<10)mm = '0'+day;



      let dateStr = "DaylyTimes-" + year+month+day;
      return dateStr;
}


function getDisplayTime(){
      let today = new Date();
      let hours= today.getHours();
      if (hours<10)mm = "0"+hours;

      let minutes = today.getMinutes();
      if (minutes<10)mm = "0"+minutes;

      let seconds = today.getSeconds();
      if (seconds<10)mm = "0"+seconds;

      let time =hours + ":" + minutes;

      console.log("timestring:" + today.toLocaleTimeString())

      return time;
}


function getMinuteTimeout(){
      let today = new Date();
      let seconds = today.getSeconds();
      let minute_timeout = (60 - seconds) * 1000;
      return minute_timeout
}

function displayTime(){
      /*Display time on HTML page*/
      getDisplayTime()
}


console.log(getDailyFileName());
console.log(getrefresh_timeout());
console.log(getDisplayTime());
console.log(getMinuteTimeout());



/*Refresh every night at 00:01*/
setInterval('autoRefresh()', getrefresh_timeout());

/*Refresh every minute*/
setInterval('displayTime()', getMinuteTimeout());
/*setInterval('displayTime()', 6000);*/