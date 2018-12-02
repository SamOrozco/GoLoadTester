import 'package:json_annotation/json_annotation.dart';

part 'Schedule.g.dart';

@JsonSerializable(nullable: false)
class Schedule {
  @JsonKey(name: "id")
  String id;
  @JsonKey(name: "name")
  String name;
  @JsonKey(name: "start_time")
  DateTime startTime;
  @JsonKey(name: "end_time")
  DateTime endTime;
  @JsonKey(name: "request_count")
  int requestCount;
  @JsonKey(name: "average_duration")
  int averageDuration;
  @JsonKey(name: "ShortestRequest")
  int shortestRequest;
  @JsonKey(name: "LongestRequest")
  int longestRequest;

  Schedule();

  factory Schedule.fromJson(json) => _$ScheduleFromJson(json);
}
