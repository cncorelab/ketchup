import * as m from 'mithril';
import * as API from 'lib/api';

export default class Route implements API.Route {
  uuid: string;
  path: string;
  file?: string;
  pageUuid?: string;

  constructor(config?: API.Route) {
    if (config) {
      this.uuid = config.uuid;
      this.path = config.path;
      this.file = config.file;
      this.pageUuid = config.pageUuid;
    }
  }

  save(): Promise<any> {
    return m.request({
      method: 'POST',
      url: '/api/v1/routes',
      data: this as API.Route
    });
  }

  static format(s: string) {
    if (s == null) {
      return '';
    }
    s = s.toLowerCase()
      .replace(/[^a-zA-Z0-9\/]+/ig, '-')
      .replace(/^-+/, '')
      .replace(/-+$/, '')
      .replace(/\/\/+/, '/');
    if (s.length > 0 && s[0] != '/') {
      s = '/' + s;
    }
    return s;
  }

  static list(): Promise<Route[]> {
    return m.request({
      method: 'GET',
      url: '/api/v1/routes'
    })
      .then((data: { routes: API.Route[] }) => {
        if (!data.routes) {
          return [];
        }
        return data.routes.map((el) => new Route(el));
      });
  }

  static saveList(routes: Route[], pageUUID: string) {
    let chain: Promise<void> = null;
    routes.map((r) => {
      r.pageUuid = pageUUID;
      if (!chain) {
        chain = r.save();
      } else {
        chain.then(() => r.save());
      }
    });
  }
}
