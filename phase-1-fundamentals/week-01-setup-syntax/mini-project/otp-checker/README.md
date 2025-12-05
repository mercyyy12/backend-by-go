# otp-checker

with pure loops, labeled break, time delays, and bufio. **NO SLICES/MAPS**

### What it does
- OTP verification
- Maximum 4 wrong attempts
- Progressive lockout delays: 5s → 15s → 30s
- Live countdown with `\r`
- Invalid input doesn't consume attempts
- Final "Device blocked!" message

### Demo
```text
Enter the OTP
111111
Wrong otp. Please wait
You have 3 attempts remaining
5...4...3...2...1...

Enter the OTP
123456
Access granted!

