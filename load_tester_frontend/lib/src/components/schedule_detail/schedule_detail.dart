import 'package:angular/angular.dart';
import 'package:angular_router/angular_router.dart';

@Component(
  selector: 'schedule-detail',
  directives: const [],
  templateUrl: 'schedule_detail.html',
  styleUrls: const [
    'schedule_detail.css',
  ],
)
class ScheduleDetailComponent {
  final Router _router;

  ScheduleDetailComponent(this._router);
}
