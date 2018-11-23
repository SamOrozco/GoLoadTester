import 'package:load_tester_frotend/src/models/models.dart';
import 'package:load_tester_frotend/src/services/services.dart';

class RequestBloc {
  final RequestService requestService;

  RequestBloc(this.requestService);

  Future<ScheduleIdResponse> createScheduleRequest(
      ScheduleRequest request) async {
    return await this.requestService.createRequest(request);
  }
}
