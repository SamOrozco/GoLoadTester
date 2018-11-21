class ScheduleRequest {
  final String url;
  final Map<String, String> headers;
  final Map<String, String> queryParams;
  final int requestCount;
  final int intervalCount;
  final String invervalType;

  ScheduleRequest(this.url, this.headers, this.queryParams, this.requestCount,
      this.intervalCount, this.invervalType);
}
