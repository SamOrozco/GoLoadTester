import 'package:json_annotation/json_annotation.dart';

part 'ScheduleIdResponse.g.dart';

@JsonSerializable(nullable: false)
class ScheduleIdResponse {
  @JsonKey(name: "schedule_id")
  String scheduleId;

  ScheduleIdResponse({this.scheduleId});

  factory ScheduleIdResponse.fromJson(Map<String, dynamic> json) =>
      _$ScheduleIdResponseFromJson(json);

  Map<String, dynamic> toJson() => _$ScheduleIdResponseToJson(this);
}
