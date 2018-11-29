import 'package:json_annotation/json_annotation.dart';

part 'schedule.g.dart';

@JsonSerializable(nullable: false)
class Schedule {
  Schedule();

  factory Schedule.fromJson(json) => _$ScheduleFromJson(json);
}
