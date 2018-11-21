import 'dart:async';

import 'package:angular/angular.dart';
import 'package:angular_components/angular_components.dart';

@Component(
  selector: 'request-url',
  directives: const [
    MaterialDropdownSelectComponent,
    MaterialInputComponent,
  ],
  templateUrl: 'request_url.html',
  styleUrls: const [
    'request_url.css',
  ],
)
class RequestUrlComponent {
  static final List<String> values = ["GET", "POST", "PUT"];
  static final SelectionOptions options = SelectionOptions.fromList(values);
  SelectionModel selectedValue = new SelectionModel.single(selected: values[0]);

  RequestUrlComponent() {}

  SelectionModel get value => selectedValue;

  void set value(SelectionModel model) {
    this.selectedValue = model;
  }

  Stream<String> get valueChanges =>
      selectedValue.selectionChanges.map((data) => data.first.toString());

  String get selectedValueString => selectedValue.selectedValues.take(1).first;
}
