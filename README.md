# Lazal

Lazal is an in-memory, minimal, fast, key-value store.

# Installation

```
docker run -p 5555:5555 sahithvibudhi/lazal
```

# Test

```
telnet 127.0.0.1 5555
> SET key1 value1
done
> GET key1
value1
```

> NOTE: Lazal is still in active development and not ready for production use.