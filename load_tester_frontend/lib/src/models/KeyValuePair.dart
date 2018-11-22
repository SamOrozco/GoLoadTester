class KeyValuePair {
  String key;
  String value;

  KeyValuePair(this.key, this.value);

  static Map<String, String> getMapFromKeyValues(List<KeyValuePair> pairs) {
    var resultMap = new Map<String, String>();
    for (var val in pairs) {
      var key = val.key;
      var value = val.value;
      if (key.isNotEmpty || value.isNotEmpty) {
        resultMap.putIfAbsent(val.key, () => val.value);
      }
    }
    return resultMap;
  }
}
