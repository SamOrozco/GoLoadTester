import 'package:angular_router/angular_router.dart';

class RoutePaths {
  static final baseRoute = RoutePath(path: '');
  static final createSchedule = RoutePath(path: 'schedule/create');
  static final scheduleDetail = RoutePath(path: 'schedule/:scheduleId');
}
