// Bluetooth Client Component
Vue.component('client-item', {
  // Template
  props: ['client'],
  template: `
  <div class="panel">
      <header>
      <h5> {{ client.Title }} </h5>
      </header>
      <section>
        <ul>
          <li>Client: {{ client.Name }}</li>
          <li>Address: {{ client.Address }} </li>
          <li>Class: {{ client.Class }} </li>
          <li>Last upated: {{ client.Timestap | updated }} </li>
        </ul>
      </section>
  </div>
  `
});

// Main Vue
var vm = new Vue({
  el: '#app',
  data: {
    clientList: []
  },
  filters: {
    updated: function (value) {
      return "Working";
    }
  }
});

setInterval(function() {
  fetch('http://192.168.0.109:3000/api/data').then(function(response) {
    return response.json();
  }).then(function(serverList) {
    vm.clientList = serverList;
  });
}, 3000);
