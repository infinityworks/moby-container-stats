# HELP container_cpu_usage_percent CPU usage percent for the specified container
# TYPE container_cpu_usage_percent gauge
container_cpu_usage_percent{container_id="84e10036a239bb738b71f4f3acb23aef97d6c2ee17172f6a91b1ce0203a3a018",container_name="epic_chebyshev"} 0
container_cpu_usage_percent{container_id="fb48eecf0f2ac58b3e68d940df3fba71c95cd4611c038087b76c3f6bdc067b17",container_name="loving_franklin"} 0.13478181366459627
# HELP container_memory_cache_bytes Current memory cache in bytes for the specified container
# TYPE container_memory_cache_bytes gauge
container_memory_cache_bytes{container_id="84e10036a239bb738b71f4f3acb23aef97d6c2ee17172f6a91b1ce0203a3a018",container_name="epic_chebyshev"} 0
container_memory_cache_bytes{container_id="fb48eecf0f2ac58b3e68d940df3fba71c95cd4611c038087b76c3f6bdc067b17",container_name="loving_franklin"} 0
# HELP container_memory_limit Memory limit as configured for the specified container
# TYPE container_memory_limit gauge
container_memory_limit{container_id="84e10036a239bb738b71f4f3acb23aef97d6c2ee17172f6a91b1ce0203a3a018",container_name="epic_chebyshev"} 2.087100416e+09
container_memory_limit{container_id="fb48eecf0f2ac58b3e68d940df3fba71c95cd4611c038087b76c3f6bdc067b17",container_name="loving_franklin"} 2.087100416e+09
# HELP container_memory_usage_bytes Current memory usage in bytes for the specified container
# TYPE container_memory_usage_bytes gauge
container_memory_usage_bytes{container_id="84e10036a239bb738b71f4f3acb23aef97d6c2ee17172f6a91b1ce0203a3a018",container_name="epic_chebyshev"} 6.43072e+06
container_memory_usage_bytes{container_id="fb48eecf0f2ac58b3e68d940df3fba71c95cd4611c038087b76c3f6bdc067b17",container_name="loving_franklin"} 2.306048e+06
# HELP container_memory_usage_percent Current memory usage percent for the specified container
# TYPE container_memory_usage_percent gauge
container_memory_usage_percent{container_id="84e10036a239bb738b71f4f3acb23aef97d6c2ee17172f6a91b1ce0203a3a018",container_name="epic_chebyshev"} 0.3081174221758193
container_memory_usage_percent{container_id="fb48eecf0f2ac58b3e68d940df3fba71c95cd4611c038087b76c3f6bdc067b17",container_name="loving_franklin"} 0.1104905150859785
# HELP container_net_rx_bytes Network RX Bytes
# TYPE container_net_rx_bytes gauge
container_net_rx_bytes{container_id="84e10036a239bb738b71f4f3acb23aef97d6c2ee17172f6a91b1ce0203a3a018",container_name="epic_chebyshev",interface="eth0"} 3855
container_net_rx_bytes{container_id="fb48eecf0f2ac58b3e68d940df3fba71c95cd4611c038087b76c3f6bdc067b17",container_name="loving_franklin",interface="eth0"} 1618
# HELP container_net_rx_dropped Network RX Dropped Packets
# TYPE container_net_rx_dropped gauge
container_net_rx_dropped{container_id="84e10036a239bb738b71f4f3acb23aef97d6c2ee17172f6a91b1ce0203a3a018",container_name="epic_chebyshev",interface="eth0"} 0
container_net_rx_dropped{container_id="fb48eecf0f2ac58b3e68d940df3fba71c95cd4611c038087b76c3f6bdc067b17",container_name="loving_franklin",interface="eth0"} 0
# HELP container_net_rx_errors Network RX Packet Errors
# TYPE container_net_rx_errors gauge
container_net_rx_errors{container_id="84e10036a239bb738b71f4f3acb23aef97d6c2ee17172f6a91b1ce0203a3a018",container_name="epic_chebyshev",interface="eth0"} 0
container_net_rx_errors{container_id="fb48eecf0f2ac58b3e68d940df3fba71c95cd4611c038087b76c3f6bdc067b17",container_name="loving_franklin",interface="eth0"} 0
# HELP container_net_rx_packets Network RX Packets
# TYPE container_net_rx_packets gauge
container_net_rx_packets{container_id="84e10036a239bb738b71f4f3acb23aef97d6c2ee17172f6a91b1ce0203a3a018",container_name="epic_chebyshev",interface="eth0"} 36
container_net_rx_packets{container_id="fb48eecf0f2ac58b3e68d940df3fba71c95cd4611c038087b76c3f6bdc067b17",container_name="loving_franklin",interface="eth0"} 23
# HELP container_net_tx_bytes Network TX Bytes
# TYPE container_net_tx_bytes gauge
container_net_tx_bytes{container_id="84e10036a239bb738b71f4f3acb23aef97d6c2ee17172f6a91b1ce0203a3a018",container_name="epic_chebyshev",interface="eth0"} 3908
container_net_tx_bytes{container_id="fb48eecf0f2ac58b3e68d940df3fba71c95cd4611c038087b76c3f6bdc067b17",container_name="loving_franklin",interface="eth0"} 0
# HELP container_net_tx_dropped Network TX Dropped Packets
# TYPE container_net_tx_dropped gauge
container_net_tx_dropped{container_id="84e10036a239bb738b71f4f3acb23aef97d6c2ee17172f6a91b1ce0203a3a018",container_name="epic_chebyshev",interface="eth0"} 0
container_net_tx_dropped{container_id="fb48eecf0f2ac58b3e68d940df3fba71c95cd4611c038087b76c3f6bdc067b17",container_name="loving_franklin",interface="eth0"} 0
# HELP container_net_tx_errors Network TX Packet Errors
# TYPE container_net_tx_errors gauge
container_net_tx_errors{container_id="84e10036a239bb738b71f4f3acb23aef97d6c2ee17172f6a91b1ce0203a3a018",container_name="epic_chebyshev",interface="eth0"} 0
container_net_tx_errors{container_id="fb48eecf0f2ac58b3e68d940df3fba71c95cd4611c038087b76c3f6bdc067b17",container_name="loving_franklin",interface="eth0"} 0
# HELP container_net_tx_packets Network TX Packets
# TYPE container_net_tx_packets gauge
container_net_tx_packets{container_id="84e10036a239bb738b71f4f3acb23aef97d6c2ee17172f6a91b1ce0203a3a018",container_name="epic_chebyshev",interface="eth0"} 24
container_net_tx_packets{container_id="fb48eecf0f2ac58b3e68d940df3fba71c95cd4611c038087b76c3f6bdc067b17",container_name="loving_franklin",interface="eth0"} 0