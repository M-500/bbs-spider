import { JSEncrypt } from 'jsencrypt'

export default {
  encrypt (data) {
    let encryptor = new JSEncrypt()
    encryptor.setPublicKey('MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDJs2456KyzZ5fLVfEMlO0VRpnFr91xgYPWfoDzB4B0LczfgGjCsJDNbVi0KdBbgMZG6JFiKaIBLPCJi/JQjrgr8YFh+H6btXgdYiBcFHYVvQ8WR5NTTMOKYw4xJIqX5CLIh5P+8F3dxlYKQ35kc5R+9jn5v8tCk4+xQJppAhYpmwIDAQAB')
    return encryptor.encrypt(data)
  }
}
