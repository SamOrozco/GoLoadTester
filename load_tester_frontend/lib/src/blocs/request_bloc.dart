import 'package:load_tester_frotend/src/models/models.dart';
import 'package:load_tester_frotend/src/services/services.dart';

class RequestBloc {
  final RequestService requestService;

  RequestBloc(this.requestService);

  void createScheduleRequest(ScheduleRequest request) {
    this.requestService.createRequest(request);
  }
}
