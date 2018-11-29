import 'package:angular/angular.dart';
import 'package:angular_router/angular_router.dart';
import 'package:firebase/firestore.dart' as fs;

@Component(
  selector: 'schedule-detail',
  directives: const [
    OnActivate,
    OnInit,
  ],
  templateUrl: 'schedule_detail.html',
  styleUrls: const [
    'schedule_detail.css',
  ],
)
class ScheduleDetailComponent implements OnActivate {
  final Router _router;
  final fs.Firestore _store;
  String _scheduleId;
  fs.DocumentReference _scheduleDoc;
  Stream<fs.QuerySnapshot> _requestStream;

  ScheduleDetailComponent(this._router, this._store);

  void handleScheduleSnapShots() {
    _scheduleDoc.onSnapshot.listen((data) => print(data.data()));
  }

  void handleRequestSnapShots() {
    _requestStream.listen((data) => print(data.docs));
  }

  @override
  // on activate we want to get the schedId from the path param
  void onActivate(_, RouterState current) async {
    String schedId = _getScheduleIdFromUrl(current.toUrl());
    this._scheduleId = schedId;
    _scheduleDoc = _store.doc("schedules/$_scheduleId");
    _requestStream = _store
        .collection("requests")
        .where("ScheduleId", "==", _scheduleId)
        .onSnapshot;
    handleRequestSnapShots();
    handleScheduleSnapShots();
  }

  String _getScheduleIdFromUrl(String url) {
    return url.split("/")[1];
  }
}
