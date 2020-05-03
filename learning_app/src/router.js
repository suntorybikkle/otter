import Vue from 'vue';
import Router from 'vue-router';
import Login from './components/login/Login.vue';
import StudyHistory from './components/history/StudyHistory.vue';
import Display from './components/main/Display.vue';


Vue.use(Router);

export default new Router({
    routes: [
        { path: '/', component: Display },
        { path: '/login', component: Login },
        { path:'/history', component: StudyHistory}
    ]
}) 