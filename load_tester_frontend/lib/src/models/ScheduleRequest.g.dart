// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'ScheduleRequest.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

ScheduleRequest _$ScheduleRequestFromJson(Map<String, dynamic> json) {
  return ScheduleRequest(
      url: json['url'] as String,
      requestType: json['requestType'] as String,
      headers: Map<String, String>.from(json['headers'] as Map),
      queryParams: Map<String, String>.from(json['queryParams'] as Map),
      requestCount: json['requestCount'] as int,
      intervalCount: json['intervalCount'] as int,
      intervalType: json['intervalType'] as String,
      name: json['name'] as String);
}

Map<String, dynamic> _$ScheduleRequestToJson(ScheduleRequest instance) =>
    <String, dynamic>{
      'url': instance.url,
      'requestType': instance.requestType,
      'headers': instance.headers,
      'queryParams': instance.queryParams,
      'requestCount': instance.requestCount,
      'intervalCount': instance.intervalCount,
      'intervalType': instance.intervalType,
      'name': instance.name
    };
