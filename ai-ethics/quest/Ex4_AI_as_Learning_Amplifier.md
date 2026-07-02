
# Exercise 4: AI as Learning Amplifier
# Phase 1: Build Foundation (you first)

Simple scenario design (5–10 devices) with justification.

# Network Devices (8 Devices)
Devices Included:

- ISP Modem
- Router
- Switch
- Wireless Access Point (WAP)
- Desktop PC
- Laptop
- Network Printer
- File Server/NAS

# Simple justification

                Internet
                    │
               [ ISP Modem ]
                    │
               [ Router ]
            (Firewall + DHCP)
                    │
               [ Switch ]
        ┌──────────┼───────────┬───────────┐
        │          │           │           │
   [Desktop]   [Printer]   [File Server] [WAP]
                                             │
                                      ┌──────┴──────┐
                                      │             │
                                   [Laptop]     (WiFi Devices)

# Explanation:
The modem connects the network to the internet, while the router manages traffic, assigns IP addresses, and provides basic security by separating the private network from the public internet. A switch allows multiple wired devices to communicate efficiently, and a wireless access point enables mobile devices to connect without cables. The desktop uses a stable wired connection for reliable work, the laptop provides mobility through Wi-Fi, the network printer allows all users to share one printing resource, and the file server (NAS) stores and backs up shared files in a central location for easy access and organization.


# Phase 2: Strategic AI Use
# Test your understanding, explore edge cases with targeted questions, then validate.

In this simple network design, key edge cases include internet or modem failure, router or DHCP malfunction, switch overload or broadcast storms, Wi-Fi congestion, storage limits on the NAS, power outages, IP address conflicts, printer queue errors, and security risks from unsegmented guest access; most of these issues arise from single points of failure (like the router or power supply) and can be mitigated with backups, proper configuration, traffic control, network segmentation, and protective hardware like a UPS.


# Phase 3: Real Application

Design a small smart-city network (1,000 IoT sensors, 50 traffic lights, 10 emergency vehicles). Decide protocols, justify choices, list failure points, then refine with AI feedback.

# Network design:

                         ┌──────────────────────────┐
                         │     City Control Center  │
                         │ (Traffic & Emergency AI) │
                         └───────────┬──────────────┘
                                     │
                              Fiber Backbone (Metro Ethernet)
                                     │
                     ┌───────────────┼────────────────┐
                     │               │                │
               [Edge Gateway]   [Edge Gateway]   [Edge Gateway]
               (District A)     (District B)     (District C)
                   │                │                │
        ┌──────────┼───────┐  ┌────┼──────┐   ┌─────┼───────┐
        │          │       │  │    │      │   │     │        │
   IoT Sensors  Traffic  Cameras  IoT Sensors  Traffic  IoT Sensors
   (LoRaWAN)     Lights  (IP)     (LoRaWAN)     Lights   (LoRaWAN)

                         ↑
               4G/5G Cellular Network
                         ↑
                Emergency Vehicles (10)
              (GPS + Video + Priority Signal)


# Protocols and Why:

| Component                 | Protocol               | Reason       |
| ------------------------- | ---------------------- | -------------------------------
| IoT Sensors               | **LoRaWAN**            | Long range (5–15 km), ultra-low power, cheap batteries (5+ years) |
| Traffic Lights            | **Ethernet + MQTT**    | Reliable, real-time control messages                              |
| Emergency Vehicles        | **4G/5G LTE + GPS**    | Mobility + wide coverage                                          |
| Gateways → Control Center | **Fiber + IP/MPLS**    | High bandwidth & low latency                                      |
| Data Messaging            | **MQTT**               | Lightweight publish/subscribe IoT messaging                       |
| Video Cameras             | **RTSP over IP**       | Real-time streaming                                               |
| Security                  | **TLS/SSL encryption** | Prevent hacking of city infrastructure                            |

# Justification of Design:
This design works because it separates responsibilities:

LoRaWAN → efficient for thousands of small sensors

Fiber → reliable city backbone

5G → mobility for vehicles

MQTT → scalable communication

Gateways → reduce network congestion


# Failure Point
1. Gateway Failure
2. Fiber Backbone Cut
3. Cellular Network Failure
4. Control Center Server Crash
5. Cyberattack
6. Power Outage
7. Sensor Battery Depletion


# Reflection:

# % human judgment vs. AI contribution
    - 50% human and 50% AI

# Could you defend decisions without AI?
    - I will defend decisions withou AI depending on the premise of such decision.

# What will you still remember in 6 months?
    -   Yes, i will still remember

# Did AI make you sharper, or think for you?
    - AI makes me sharper.


