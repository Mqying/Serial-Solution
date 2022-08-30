export default [
  {
    path: '/user',
    layout: false,
    routes: [
      {
        name: 'login',
        path: '/user/login',
        component: './User/Login',
      },
      {
        component: './404',
      },
    ],
  },
  {
    path: '/introduction',
    name: 'introduction',
    component: './Introduction',
    access: 'introduction'
  },
  {
    path: '/flash',
    name: 'query',
    icon: 'dashboard',
    component: './FlashQuery',
    access: 'flash'
  },
  {
    path: '/dielectric',
    name: 'query',
    icon: 'dashboard',
    component: './DielectricQuery',
    access: 'dielectric'
  },
  {
    path: '/water',
    name: 'query',
    icon: 'dashboard',
    component: './WaterQuery',
    access: 'water'
  },
  {
    path: '/acid',
    name: 'query',
    icon: 'dashboard',
    component: './AcidQuery',
    access: 'acid'
  },
  {
    path: '/history',
    name: 'history',
    icon: 'table',
    component: './History',
  },
  {
    path: '/',
    redirect: '/history',
  },
  {
    component: './404',
  },
];
