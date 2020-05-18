# エニグマ暗号

## 使い方

### 1. エニグマを初期化する
引数はローターの数
```go:
machine := initEnigma(rotorAmount)
```

### 2. 暗号化する
引数は暗号化したい文字列
```go:
encrypt := machine.Encrypt(text)
```

### 3. 復号化する
引数は復号化したい文字列
```go:
decrypt := machine.Decrypt(text)
```
