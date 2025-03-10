# Modular Arithmetic and Prime Numbers in Cryptography

## 1. **Modular Arithmetic and Prime Numbers in Cryptography**

### 1.1 Modular Arithmetic
- **Definition**: Modular arithmetic is a system of arithmetic for integers where numbers "wrap around" after reaching a certain value (the modulus).
    - Example: `a ≡ b mod m` means that `a` and `b` have the same remainder when divided by `m`.
- **Applications**:
    - **Modular Exponentiation**: Used in RSA and DSA for encryption and decryption.
        - Example: `c ≡ m^e mod n`
    - **Modular Inverses**: Essential for key generation in RSA.
        - Example: `d ≡ e^(-1) mod ϕ(n)`

### 1.2 Prime Numbers
- **Definition**: A prime number is a natural number greater than 1 that has no positive divisors other than 1 and itself.
- **Importance in Cryptography**:
    - **RSA Security**: Relies on the difficulty of factoring the product of two large prime numbers.
    - **Key Generation**: Large primes are used to generate public and private keys.

---

## 2. **Symmetric vs. Asymmetric Key Cryptography**

### 2.1 Symmetric Key Cryptography
- **Definition**: Uses the same key for both encryption and decryption.
- **Examples**: AES, DES, 3DES.
- **Advantages**:
    - Faster than asymmetric cryptography.
    - Suitable for encrypting large amounts of data.
- **Disadvantages**:
    - Key distribution is a challenge.
    - Not scalable for large networks.

### 2.2 Asymmetric Key Cryptography
- **Definition**: Uses a pair of keys (public and private) for encryption and decryption.
- **Examples**: RSA, DSA, ECC.
- **Advantages**:
    - Solves the key distribution problem.
    - Provides digital signatures and authentication.
- **Disadvantages**:
    - Slower than symmetric cryptography.
    - Requires more computational resources.

---

## 3. **RSA Algorithm**

### 3.1 Key Generation
1. Choose two large prime numbers `p` and `q`.
2. Compute `n = p * q` and `ϕ(n) = (p-1)(q-1)`.
3. Choose an integer `e` such that `1 < e < ϕ(n)` and `gcd(e, ϕ(n)) = 1`.
4. Compute `d` as the modular inverse of `e`:
    - `d ≡ e^(-1) mod ϕ(n)`
5. Public key: `(e, n)`, Private key: `(d, n)`.

### 3.2 Encryption
- To encrypt a message `m`:
    - `c ≡ m^e mod n`

### 3.3 Decryption
- To decrypt ciphertext `c`:
    - `m ≡ c^d mod n`

### 3.4 Attacks on RSA
- **Brute Force**: Trying all possible private keys (Specially when `n` is small or using Quantum machines).
- **Factorization Attacks**: Factoring `n` into `p` and `q`.
- **Timing Attacks**: Exploiting the time taken for decryption.

### 3.5 Usage of RSA

- SSH: `ssh-keygen -t rsa -b 2048 -f my_rsa_key`.
- PEM Format: `openssl rsa -in my_rsa_key -pubout -outform PEM -out my_rsa_key.pub`.

---

## 4. **Diffie-Hellman Key Exchange**

### 4.1 Key Exchange Protocol
1. Two parties agree on a large prime `p` and a base `g` (primitive root modulo `p`).
2. Each party chooses a private key:
    - Alice: `a`, Bob: `b`.
3. Each party computes their public key:
    - Alice: `A ≡ g^a mod p`, Bob: `B ≡ g^b mod p`.
4. They exchange public keys and compute the shared secret:
    - Alice: `S ≡ B^a mod p`, Bob: `S ≡ A^b mod p`.

### 4.2 Discrete Logarithm Problem
- The security of Diffie-Hellman relies on the difficulty of solving the discrete logarithm problem:
    - Given `g`, `p`, and `A ≡ g^a mod p`, find `a`.

---

## **Conclusion**

- **Modular Arithmetic** and **Prime Numbers** are fundamental to modern cryptography, especially in RSA and DSA algorithms.
- **Symmetric Cryptography** is efficient for bulk data encryption, while **Asymmetric Cryptography** solves key distribution issues.
- **RSA** relies on the difficulty of factoring large numbers, and **Diffie-Hellman** relies on the discrete logarithm problem.
- The security of these algorithms depends on the computational infeasibility of solving these mathematical problems.