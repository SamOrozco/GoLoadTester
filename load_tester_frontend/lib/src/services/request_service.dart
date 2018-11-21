import 'package:http/http.dart';
import 'package:load_tester_frotend/src/models/models.dart';
import 'dart:convert';

class RequestService {
  final Client _http;

  RequestService(this._http);
  void createRequest(ScheduleRequest request) {
    var jsonString = jsonEncode(request);
    print(jsonString);
  }
}
