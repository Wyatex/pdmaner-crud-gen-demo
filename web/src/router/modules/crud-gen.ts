const crudGen: AuthRoute.Route = {
  name: 'crud-gen',
  path: '/crud-gen',
  component: 'self',
  meta: {
    title: 'CRUD生成示例',
    i18nTitle: 'routes.crud-gen._value',
    requiresAuth: true,
    keepAlive: true,
    singleLayout: 'basic',
    permissions: ['super', 'admin', 'user'],
    icon: 'cib:app-store',
    order: 10
  }
};

export default crudGen;
