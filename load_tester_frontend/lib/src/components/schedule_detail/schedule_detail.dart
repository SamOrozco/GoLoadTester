import 'package:angular/angular.dart';
import 'package:angular_components/angular_components.dart';
import 'package:angular_router/angular_router.dart';
import 'package:firebase/firestore.dart' as fs;
import 'package:load_tester_frotend/src/components/components.dart';
import 'package:load_tester_frotend/src/models/Schedule.dart';
import 'package:load_tester_frotend/src/models/models.dart';

@Component(
  selector: 'schedule-detail',
  directives: const [
    OnActivate,
    OnInit,
    NgModel,
    NgFor,
    NgIf,
    RequestResponseComponent,
    MaterialIconComponent,
    coreDirectives,
    LoadingComponent,
  ],
  templateUrl: 'schedule_detail.html',
  styleUrls: const [
    'schedule_detail.css',
    'package:angular_components/css/mdc_web/card/mdc-card.scss.css',
  ],
)
class ScheduleDetailComponent implements OnActivate {
  final Router _router;
  final fs.Firestore _store;
  String _scheduleId;
  fs.DocumentReference _scheduleDoc;
  Stream<fs.QuerySnapshot> _requestStream;
  Schedule schedule = new Schedule();
  List<RequestResponse> responses = new List<RequestResponse>();

  // loading varaibles to help with the chopiness
  bool schedLoaded = false;
  bool reqLoaded = false;

  ScheduleDetailComponent(this._router, this._store);

  void handleScheduleSnapShots() async {
    var snap = await _scheduleDoc.get();
    _scheduleDoc.onSnapshot.listen((data) {
      if (data.exists) {
        this.schedule = Schedule.fromJson(data.data());
        this.schedLoaded = true;
      }
    });
  }

  void handleRequestSnapShots() {
    _requestStream.listen((data) {
      for (var doc in data.docChanges()) {
        RequestResponse resp = RequestResponse.fromJson(doc.doc.data());
        this.responses.add(resp);
        this.reqLoaded = true;
      }
    });
  }

  @override
  // on activate we want to get the schedId from the path param
  void onActivate(_, RouterState current) async {
    this.schedLoaded = false;
    this.reqLoaded = false;
    String schedId = _getScheduleIdFromUrl(current.toUrl());
    this._scheduleId = schedId;
    print(this._scheduleId);
    // get our doc ref
    _scheduleDoc = _store.doc("schedules/$_scheduleId");
    // get a stream of updates from our schedule requests
    _requestStream = _store
        .collection("requests")
        .where("schedule_id", "==", _scheduleId)
        .onSnapshot;

    handleRequestSnapShots();
    handleScheduleSnapShots();
  }

  String _getScheduleIdFromUrl(String url) {
    return url.split("/")[1];
  }

  void handleDocChange(fs.DocumentChange change) {
    var type = change.type;
    switch (type) {
      case "added":
        break;
      case "removed":
        break;
      case "modified":
        break;
    }
  }
}
