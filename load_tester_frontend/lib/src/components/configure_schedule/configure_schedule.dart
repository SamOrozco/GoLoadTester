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
  static final SelectionOptions options =
      SelectionOptions.fromList(timeUnitValues);
  SelectionModel selectedValue =
      new SelectionModel.single(selected: timeUnitValues[0]);
  static final List<String> timeUnitValues = [
    "Milliseconds",
    "Seconds",
    "Minutes",
  ];

  String requestCountString;
  String timeIntervalString;
  String scheduleName;

  int get requestCount => int.tryParse(requestCountString);

  int get timeInterval => int.tryParse(timeIntervalString);

  ConfigureScheduleComponent();

  SelectionModel get value => selectedValue;

  void set value(SelectionModel model) {
    this.selectedValue = model;
  }

  String get selectedTimeTypeString =>
      selectedValue.selectedValues.take(1).first;
}
