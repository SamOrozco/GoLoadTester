import 'package:angular/angular.dart';
import 'package:http/browser_client.dart';
import 'package:load_tester_frotend/src/services/services.dart';
import 'package:load_tester_frotend/src/models/models.dart';
import 'package:load_tester_frotend/app_component.template.dart' as ng;

// ignore: uri_has_not_been_generated
import 'main.template.dart' as mainGenerated;

RequestService _requestService;

RequestService requestServiceFactory() => _requestService;

@GenerateInjector(
  const [
    const FactoryProvider(RequestService, requestServiceFactory),
  ], // You can use routerProviders in production
)
final InjectorFactory appInjector = mainGenerated.appInjector$Injector;

void main() {
  var client = new BrowserClient();
  var environment = new Environment("http://connector1.ngrok.io");
  _requestService = new RequestService(client, environment);
  runApp(ng.AppComponentNgFactory, createInjector: appInjector);
}
