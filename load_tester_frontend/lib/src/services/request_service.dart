import 'dart:convert';

import 'package:http/http.dart';
import 'package:load_tester_frotend/src/models/models.dart';

class RequestService {
  final Client _http;
  final Environment _environment;

  RequestService(this._http, this._environment);

  Future<ScheduleIdResponse> createRequest(ScheduleRequest request) async {
    var body = jsonEncode(request.toJson());
    var resp = await _http.post(
        "https://gentle-wave-92777.herokuapp.com/schedule/request",
        headers: {"Content-type": "application/json"},
        body: body);
    return ScheduleIdResponse.fromJson(jsonDecode(resp.body));
  }
}
