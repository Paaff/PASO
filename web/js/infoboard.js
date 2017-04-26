// Bluetooth Client Component
Vue.component('client-item', {
  // Template
  props: ['client'],
  template: `
  <div class="panel">
      <header>
      <h5> {{ client.title }} </h5>
      </header>
      <section>
        <ul>
          <li>Client: {{ client.name }}</li>
          <li>Address: {{ client.address }} </li>
          <li>Class: {{ client.class }} </li>
          <li>Timestamp: {{client.timestamp }} </li>
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
  }
});

setInterval(function() {
  fetch('http://localhost:3000/api/data').then(function(response) {
    return response.json();
  }).then(function(client) {
    vm.clientList = client;
  });
}, 3000);
