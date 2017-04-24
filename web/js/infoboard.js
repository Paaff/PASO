
// Bluetooth Client Component
Vue.component('client-item', {
  props: ['client'],
  template: ''
});

// Main Vue
new Vue({
  el: '#app',
  data: {
    clientList: [
      { name: 'Mathias', btaddress: '12:5A:3F:3D:15', class: 'Smartphone class', timestamp: '13:37'},
      { name: 'Peter', btaddress: '3B:2C:5F:7E:12', class: 'Smartphone class', timestamp: '01:37'}
    ]
  }
});
