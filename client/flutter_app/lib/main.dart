import 'package:flutter/material.dart';
import 'package:web_socket_channel/web_socket_channel.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return const MaterialApp(
      home: MyHomePage(),
    );
  }
}

class MyHomePage extends StatefulWidget {
  const MyHomePage({super.key});

  @override
  State<MyHomePage> createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  final _channel = WebSocketChannel.connect(
    Uri.parse('ws://localhost:8080/ws/run/demo-run'),
  );

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('PEA'),
      ),
      body: Row(
        children: [
          // Conversation list skeleton
          Container(
            width: 200,
            color: Colors.grey[200],
            child: const Center(
              child: Text('Conversations'),
            ),
          ),
          // Main content area
          Expanded(
            child: Column(
              children: [
                // Prompt library skeleton
                Container(
                  height: 100,
                  color: Colors.grey[300],
                  child: const Center(
                    child: Text('Prompt Library'),
                  ),
                ),
                // Streaming token display
                Expanded(
                  child: StreamBuilder(
                    stream: _channel.stream,
                    builder: (context, snapshot) {
                      return Text(snapshot.hasData ? '${snapshot.data}' : '');
                    },
                  ),
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }

  @override
  void dispose() {
    _channel.sink.close();
    super.dispose();
  }
}
