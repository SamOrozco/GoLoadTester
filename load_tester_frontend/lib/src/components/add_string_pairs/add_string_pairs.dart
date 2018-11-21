import 'dart:collection';

import 'package:angular/angular.dart';
import 'package:angular_components/angular_components.dart';
import 'package:angular_forms/angular_forms.dart';
import 'package:load_tester_frotend/src/models/KeyValuePair.dart';

@Component(
  selector: 'add-string-pairs',
  directives: const [
    materialInputDirectives,
    MaterialButtonComponent,
    MaterialIconComponent,
    NgFor,
    NgModel,
    formDirectives,
  ],
  templateUrl: 'add_string_pairs.html',
  styleUrls: const [
    'add_string_pairs.css',
  ],
)
class AddStringPairsComponent {
  List<KeyValuePair> keyValuePairs = new List<KeyValuePair>();

  // set name of this puppy
  String name = "Add String Pair Component";
  String keyTitle = "Key";
  String valueTitle = "Value";

  AddStringPairsComponent() {
    appendEmptyPair();
  }

  void addNewItem(KeyValuePair pair) {
    var idx = this.keyValuePairs?.indexOf(pair);
    // if this is last item in list
    if (idx == this.keyValuePairs.length - 1) {
      appendEmptyPair();
    }
  }

  void appendEmptyPair() {
    this.addPair(new KeyValuePair("", ""));
  }

  void removePair(KeyValuePair pair) {
    if (this.keyValuePairs?.length == 1) {
      this.keyValuePairs?.remove(pair);
      appendEmptyPair();
    } else {
      this.keyValuePairs?.remove(pair);
    }
  }

  void addPair(KeyValuePair pair) {
    this.keyValuePairs?.add(pair);
  }
}
