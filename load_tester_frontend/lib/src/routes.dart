import 'package:angular_router/angular_router.dart';

import 'package:load_tester_frotend/src/components/schedule_detail/schedule_detail.template.dart'
    as det;

import 'package:load_tester_frotend/src/components/add_request/add_request.template.dart'
    as add;

import 'route_paths.dart';

export 'route_paths.dart';

class Routes {
  static final addRequestRoute = RouteDefinition(
    routePath: RoutePaths.createSchedule,
    component: add.AddRequestNgFactory,
  );

  static final baseRoute = RouteDefinition(
    routePath: RoutePaths.baseRoute,
    component: add.AddRequestNgFactory,
  );

  static final scheduleDetailRoute = RouteDefinition(
    routePath: RoutePaths.scheduleDetail,
    component: det.ScheduleDetailComponentNgFactory,
  );

  static final all = <RouteDefinition>[
    addRequestRoute,
    baseRoute,
    scheduleDetailRoute,
  ];
}
