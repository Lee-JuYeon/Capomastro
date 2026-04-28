---
ide: terminal
appName: Terminal
description: ROS 2 Robotics Package (Python + ament_cmake)
---

ros2 pkg create generates a ROS 2 ament_cmake package with Python node scaffold.
Requires ROS 2 Jazzy (or Humble) installed and sourced. Build: colcon build.

```json
[
  { "action": "shell", "cmd": "ros2 pkg create {name} --build-type ament_cmake --node-name {name}_node --dependencies rclpy std_msgs", "cwd": "{outputDir}", "timeout": 15000 }
]
```
