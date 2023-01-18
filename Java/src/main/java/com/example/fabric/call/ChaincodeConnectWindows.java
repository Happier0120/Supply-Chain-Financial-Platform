package com.example.fabric.call;

import com.example.fabric.constant.Submit;
import org.bouncycastle.util.io.pem.PemReader;
import org.hyperledger.fabric.gateway.*;

import java.io.FileReader;
import java.io.IOException;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.security.InvalidKeyException;
import java.security.PrivateKey;
import java.security.cert.CertificateException;
import java.security.cert.X509Certificate;


public class ChaincodeConnectWindows {

    static {
        System.setProperty("org.hyperledger.fabric.sdk.service_discovery.as_localhost", "true");
    }

    // helper function for getting connected to the gateway
    public static Gateway connect() throws Exception {
        // 创建钱包，加载用户密钥
        Path walletPath = Paths.get("D:\\学校\\中交智运\\实现\\data\\wallet");
        Wallet wallet = Wallets.newFileSystemWallet(walletPath);
        populateWallet(wallet);
        // 网络连接，加载fabric信息与peer节点密钥
        Path networkConfigPath = Paths.get("D:\\学校\\中交智运\\实现\\data\\connection.yaml");
        // 连接建立，后续可用于执行链码
        Gateway.Builder builder = Gateway.createBuilder();
        builder.identity(wallet, "appUser").networkConfig(networkConfigPath).discovery(true);
        return builder.connect();
    }

    public static void populateWallet(Wallet wallet) throws InvalidKeyException, CertificateException, IOException {
        String certPathAdmin = "D:\\学校\\中交智运\\实现\\data\\peerOrganizations\\" + Submit.users.get(Submit.currentUser) + ".blockchain.com\\ca\\ca." + Submit.users.get(Submit.currentUser) + ".blockchain.com-cert.pem";
        String keyPathAdmin = "D:\\学校\\中交智运\\实现\\data\\peerOrganizations\\" + Submit.users.get(Submit.currentUser) + ".blockchain.com\\ca\\priv_sk";
        X509Certificate certAdmin = Identities.readX509Certificate(new PemReader(new FileReader(certPathAdmin)));
        PrivateKey keyAdmin = Identities.readPrivateKey(new PemReader(new FileReader(keyPathAdmin)));
        X509Identity identityAdmin = Identities.newX509Identity(Submit.users.get(Submit.currentUser), certAdmin, keyAdmin);
        wallet.put("admin", identityAdmin);

        String certPath = "D:\\学校\\中交智运\\实现\\data\\peerOrganizations\\" + Submit.users.get(Submit.currentUser) + ".blockchain.com\\users\\User1@" + Submit.users.get(Submit.currentUser) + ".blockchain.com\\msp\\signcerts\\User1@" + Submit.users.get(Submit.currentUser) + ".blockchain.com-cert.pem";
        String keyPath = "D:\\学校\\中交智运\\实现\\data\\peerOrganizations\\" + Submit.users.get(Submit.currentUser) + ".blockchain.com\\users\\User1@" + Submit.users.get(Submit.currentUser) + ".blockchain.com\\msp\\keystore\\priv_sk";
        X509Certificate cert = Identities.readX509Certificate(new PemReader(new FileReader(certPath)));
        PrivateKey key = Identities.readPrivateKey(new PemReader(new FileReader(keyPath)));
        X509Identity identity = Identities.newX509Identity(Submit.users.get(Submit.currentUser), cert, key);
        wallet.put("appUser", identity);
    }
}
