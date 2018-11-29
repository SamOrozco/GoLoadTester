import 'package:angular/angular.dart';
import 'package:angular_router/angular_router.dart';
import 'package:http/browser_client.dart';
import 'package:load_tester_frotend/src/services/services.dart';
import 'package:load_tester_frotend/src/models/models.dart';
import 'package:load_tester_frotend/app_component.template.dart' as ng;
import 'package:firebase/firebase.dart';
import 'package:firebase/firestore.dart' as fs;

// ignore: uri_has_not_been_generated
import 'main.template.dart' as mainGenerated;

RequestService _requestService;

RequestService requestServiceFactory() => _requestService;

fs.Firestore _store;

fs.Firestore campaignServiceFactory() => _store;

@GenerateInjector(
  const [
    routerProvidersHash,
    const FactoryProvider(RequestService, requestServiceFactory),
    const FactoryProvider(fs.Firestore, campaignServiceFactory),
  ], // You can use routerProviders in production
)
final InjectorFactory appInjector = mainGenerated.appInjector$Injector;

void main() {
  initializeApp(
      apiKey: "AIzaSyBk4DpPIx-hnHrMkGUDZkffJMrOSK9c588",
      authDomain: "load-tester-orozco.firebaseapp.com",
      databaseURL: "load-tester-orozco.firebaseio.com",
      projectId: "load-tester-orozco",
      storageBucket: "load-tester-orozco.appspot.com");
  _store = firestore();
  var client = new BrowserClient();
  var environment = new Environment("http://connector1.ngrok.io");
  _requestService = new RequestService(client, environment);
  runApp(ng.AppComponentNgFactory, createInjector: appInjector);
}
