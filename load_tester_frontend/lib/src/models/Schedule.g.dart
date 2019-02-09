// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'Schedule.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

Schedule _$ScheduleFromJson(Map<String, dynamic> json) {
  return Schedule()
    ..id = json['id'] as String
    ..name = json['name'] as String
    ..startTime = json['start_time'] as int
    ..endTime = json['end_time'] as int
    ..requestCount = json['request_count'] as int
    ..averageDuration = json['average_duration'] as int
    ..shortestDuration = json['shortest_duration'] as int
    ..longestDuration = json['longest_duration'] as int
    ..currentRequestCount = json['current_request_count'] as int
    ..totalDuration = json['total_duration'] as int
    ..url = json['url'] as String;
}

Map<String, dynamic> _$ScheduleToJson(Schedule instance) => <String, dynamic>{
      'id': instance.id,
      'name': instance.name,
      'start_time': instance.startTime,
      'end_time': instance.endTime,
      'request_count': instance.requestCount,
      'average_duration': instance.averageDuration,
      'shortest_duration': instance.shortestDuration,
      'longest_duration': instance.longestDuration,
      'current_request_count': instance.currentRequestCount,
      'total_duration': instance.totalDuration,
      'url': instance.url
    };
