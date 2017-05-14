// Bluetooth Client Component
Vue.component('client-item', {
  // Template
  props: ['client'],
  template: `
  <div class="panel">
      <section>
        <ul>
          <li>Address: {{ client.Address }} </li>
          <li>Class: {{ client.Class }} </li>
          <li>Last updated: {{ client.Timestamp }} </li>
        </ul>
      </section>
  </div>
  `
});

// Project Component
Vue.component('project-item', {
  // Template
  props: ['project'],
  template: `
  <div class="panel">
      <header>
      <h5> {{ project.ProjectName }} </h5>
      </header>
      <section>
        <ul>
          <li>Content: {{ project.Content }} </li>
          <li>Members:
            <div v-for="item in project.Members">
              <div> {{ item }} </div>
            </div>
          </li>
          <li>Required Permissions:
            <div v-for="item in project.RequiredPermissions">
              <div> {{ item }} </div>
            </div>
          </li>
        </ul>
      </section>
  </div>
  `
});



// Main Vue
var vm = new Vue({
  el: '#app',
  data: {
    projectList: [],
    clientList: []
  }
});

setInterval(function() {
  fetch('http://192.168.0.109:3000/api/projects').then(function(response) {
    return response.json();
  }).then(function(projectListFromServer) {
    vm.projectList = projectListFromServer;
  });
}, 3000);

setInterval(function() {
  fetch('http://192.168.0.109:3000/api/data').then(function(response) {
    return response.json();
  }).then(function(clientListFromServer) {
    vm.clientList = clientListFromServer;
  });
}, 3000);
