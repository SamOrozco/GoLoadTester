import 'package:http/http.dart';
import 'package:load_tester_frotend/src/models/models.dart';
import 'dart:convert';

class RequestService {
  final Client _http;
  final Environment _environment;

  RequestService(this._http, this._environment);

  void createRequest(ScheduleRequest request) async {
    var body = jsonEncode(request.toJson());
    print(body);
    var resp = await _http.post("https://connector1.ngrok.io/schedule/request",
        headers: {"Content-type": "application/json"}, body: body);
    print(resp.body);
  }
}
