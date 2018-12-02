import 'package:json_annotation/json_annotation.dart';

part 'RequestResponse.g.dart';

@JsonSerializable(nullable: false)
class RequestResponse {
  @JsonKey(name: "id")
  String id;
  @JsonKey(name: "schedule_id")
  String scheduleId;
  @JsonKey(name: "err_string")
  String errString;
  @JsonKey(name: "duration")
  int duration;
  @JsonKey(name: "request_url")
  String requestUrl;
  @JsonKey(name: "request_type")
  String requestType;
  @JsonKey(name: "message")
  String message;

  RequestResponse();

  factory RequestResponse.fromJson(json) => _$RequestResponseFromJson(json);
}
