import 'dart:html';
import 'package:md_toast/md_toast.dart';
import 'package:angular/angular.dart';
import 'package:angular_components/angular_components.dart';

@Component(
  selector: 'show-code',
  directives: const [
    MaterialIconComponent,
    MaterialButtonComponent,
    MdToast,
  ],
  templateUrl: 'show_code.html',
  styleUrls: const [
    'show_code.css',
  ],
)
class ShowCodeComponent {
  @Input()
  String text;
  @ViewChild('showToast')
  MdToast toastElement;

  ShowCodeComponent();

  void copyMessage() async {
    _copyToClipboard(this.text);
    toastElement.color = "#ff7f24";
    toastElement.showToast("Copied");
  }

  void _copyToClipboard(String text) {
    final textarea = new TextAreaElement();
    document.body.append(textarea);
    textarea.style.border = '0';
    textarea.style.margin = '0';
    textarea.style.padding = '0';
    textarea.style.opacity = '0';
    textarea.style.position = 'absolute';
    textarea.readOnly = true;
    textarea.value = text;
    textarea.select();
    final result = document.execCommand('copy');
    textarea.remove();
  }
}
