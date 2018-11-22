import 'package:json_annotation/json_annotation.dart';

part 'ScheduleRequest.g.dart';

@JsonSerializable(nullable: false)
class ScheduleRequest {
  @JsonKey(name: 'url')
  final String url;
  final String requestType;
  final Map<String, String> headers;
  final Map<String, String> queryParams;
  final int requestCount;
  final int intervalCount;
  final String intervalType;

  ScheduleRequest({
    this.url,
    this.requestType,
    this.headers,
    this.queryParams,
    this.requestCount,
    this.intervalCount,
    this.intervalType,
  });

  factory ScheduleRequest.fromJson(Map<String, dynamic> json) =>
      _$ScheduleRequestFromJson(json);

  Map<String, dynamic> toJson() => _$ScheduleRequestToJson(this);
}
