import 'package:angular/angular.dart';
import 'package:angular_components/angular_components.dart';

@Component(
  selector: 'configure-schedule',
  directives: const [
    materialInputDirectives,
    MaterialDropdownSelectComponent,
    MaterialInputComponent,
  ],
  templateUrl: 'configure_schedule.html',
  styleUrls: const [
    'configure_schedule.css',
  ],
)
class ConfigureScheduleComponent {
  static final List<String> timeUnitValues = [
    "Milliseconds",
    "Seconds",
    "Minutes",
  ];
  static final SelectionOptions options =
      SelectionOptions.fromList(timeUnitValues);
  SelectionModel selectedValue =
      new SelectionModel.single(selected: timeUnitValues[0]);
  ConfigureScheduleComponent();

  SelectionModel get value => selectedValue;
  void set value(SelectionModel model) {
    this.selectedValue = model;
  }

  Stream<String> get valueChanges =>
      selectedValue.selectionChanges.map((data) => data.first.toString());

  String get selectedValueString => selectedValue.selectedValues.take(1).first;
}
