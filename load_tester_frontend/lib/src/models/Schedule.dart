import 'package:intl/intl.dart';
import 'package:json_annotation/json_annotation.dart';

part 'Schedule.g.dart';

@JsonSerializable(nullable: false)
class Schedule {
  @JsonKey(name: "id")
  String id;
  @JsonKey(name: "name")
  String name;
  @JsonKey(name: "start_time")
  int startTime;
  @JsonKey(name: "end_time")
  int endTime;
  @JsonKey(name: "request_count")
  int requestCount;
  @JsonKey(name: "average_duration")
  int averageDuration;
  @JsonKey(name: "shortest_duration")
  int shortestDuration;
  @JsonKey(name: "longest_duration")
  int longestDuration;
  @JsonKey(name: "current_request_count")
  int currentRequestCount;
  @JsonKey(name: "total_duration")
  int totalDuration;
  @JsonKey(name: "url")
  String url;

  String get startTimeString {
    if (startTime == null) {
      return "";
    }
    var formatter = new DateFormat('MM/dd/yyyy HH:mm:ss');
    DateTime date = DateTime.fromMillisecondsSinceEpoch(startTime * 1000);
    return formatter.format(date);
  }

  String get endTimeString {
    if (endTime == null) {
      return "";
    }
    var formatter = new DateFormat('MM/dd/yyyy HH:mm:ss');
    DateTime date = DateTime.fromMillisecondsSinceEpoch(endTime * 1000);
    return formatter.format(date);
  }

  String get averageDurationString {
    if (averageDuration == null) {
      return "";
    }
    var secDur = averageDuration;
    return "$secDur (ms)";
  }

  String get totalDurationString {
    if (totalDuration == null) {
      return "";
    }
    var secDur = totalDuration / 1000;
    return "$secDur (ms)";
  }

  Schedule();

  factory Schedule.fromJson(json) => _$ScheduleFromJson(json);
}
