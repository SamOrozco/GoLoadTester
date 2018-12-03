import 'package:angular/angular.dart';

@Component(
  selector: 'show-code',
  directives: const [],
  templateUrl: 'show_code.html',
  styleUrls: const [
    'show_code.css',
  ],
)
class ShowCodeComponent {
  String text;

  ShowCodeComponent();
}
