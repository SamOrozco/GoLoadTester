import 'package:angular_router/angular_router.dart';

import 'package:load_tester_frotend/app_component.template.dart' as app;

//import 'package:load_tester_frotend/add_request.template.dart' as add;
import 'route_paths.dart';

export 'route_paths.dart';

class Routes {
  static final baseRoute = RouteDefinition(
    routePath: RoutePaths.createSchedule,
    component: app.AppComponentNgFactory,
  );

//  static final addRequestRoute = RouteDefinition(
//    routePath: RoutePaths.createSchedule,
//    component: add.AddRequestComponentNgFactory,
//  );

  static final all = <RouteDefinition>[
    baseRoute,
//    addRequestRoute,
  ];
}
