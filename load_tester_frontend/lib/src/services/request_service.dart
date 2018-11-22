import 'package:http/http.dart';
import 'package:load_tester_frotend/src/models/models.dart';
import 'dart:convert';

class RequestService {
  final Client _http;
  final Environment _environment;

  RequestService(this._http, this._environment);
  void createRequest(ScheduleRequest request) {
    print(json.encode(request.toJson()));
  }
}
