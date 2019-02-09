// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'RequestResponse.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

RequestResponse _$RequestResponseFromJson(Map<String, dynamic> json) {
  return RequestResponse()
    ..id = json['id'] as String
    ..scheduleId = json['schedule_id'] as String
    ..errString = json['err_string'] as String
    ..duration = json['duration'] as int
    ..requestUrl = json['request_url'] as String
    ..requestType = json['request_type'] as String
    ..message = json['message'] as String;
}

Map<String, dynamic> _$RequestResponseToJson(RequestResponse instance) =>
    <String, dynamic>{
      'id': instance.id,
      'schedule_id': instance.scheduleId,
      'err_string': instance.errString,
      'duration': instance.duration,
      'request_url': instance.requestUrl,
      'request_type': instance.requestType,
      'message': instance.message
    };
