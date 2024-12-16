FROM scylladb/scylla@sha256:1de6d5303d348896a5eab27f77b7124dcfacfd32fe9e8d3823e80d1ef3030838
# Set environment variables for ScyllaDB
ENV SCYLLA_CONFIG_DIR=/etc/scylla
ENV SCYLLA_DATA_DIR=/var/lib/scylla
ENV SCYLLA_LOG_DIR=/var/log/scylla

# Expose necessary ports
EXPOSE 7000 7001 9042 9160 10000

# Set the default command to run ScyllaDB
CMD ["scylla"]
