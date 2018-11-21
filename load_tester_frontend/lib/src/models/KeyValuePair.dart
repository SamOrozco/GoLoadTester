class KeyValuePair {
  String key;
  String value;

  KeyValuePair(this.key, this.value);

  static Map<String, String> getMapFromKeyValues(List<KeyValuePair> pairs) {
    var resultMap = new Map();
    for (var val in pairs) {
      resultMap.putIfAbsent(val.key, () => val.value);
    }
  }
}
