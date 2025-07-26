# Local HTTP File Server

Um servidor de arquivos local minimalista para compartilhamento rápido via navegador ou QR Code.


<img src="https://github.com/cristoferluch/assets/blob/main/http-fileserver.png?raw=true" width="300">

### Como Usar

```
http-fileserver.exe -f C:\Users -p 8080
```

### Opções

| Flag          | Descrição                 | Padrão      |
|---------------|---------------------------|-------------|
| `-f --folder` | Pasta a ser compartilhada | `.` (atual) |
| `-p, --port`  | Porta do servidor	        | . `8080`    |

### Recursos

- Acesso via navegador ou QR Code
- Sem cache (arquivos sempre atualizados)
- Suporte a navegação em diretórios
- Health check em `/health`