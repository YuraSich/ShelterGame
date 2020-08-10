



// material design icon
// import 'material-design-icons-iconfont/dist/material-design-icons.css'
//import '@mdi/font/css/materialdesignicons.css' // Ensure you are using css-loader
// import { library } from '@fortawesome/fontawesome-svg-core'
// import { faSpinner } from '@fortawesome/free-solid-svg-icons'
// import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
// library.add(faSpinner);
// Vue.component('font-awesome-icon', FontAwesomeIcon);



// import '@mdi/font/css/materialdesignicons.css' // Ensure you are using css-loader
import '@fortawesome/fontawesome-free/js/all.js'
// import '@fortawesome/fontawesome-free/css/all.css'
import "./components/App/_App.scss"

// import "@mdi/font/css/materialdesignicons.css"
// import 'material-design-icons-iconfont/dist/material-design-icons.css' // Ensure you are using css-loader

import Vue from 'vue'
import App from "./components/App/App";
import vuetify from "../plugins/vuetify";

new Vue({
  el: '#app',
  vuetify,
  render: h => h(App)
});
