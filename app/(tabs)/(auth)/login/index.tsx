import Input from '@/components/ui/input';
import { StyleSheet, View, Text } from 'react-native';
import CustomButton from '@/components/ui/button';
import { LinearGradient } from 'expo-linear-gradient';

export default function LoginScreen() {
  return (
    <View style={styles.container}>
      <LinearGradient
        colors={[
          'rgba(132, 34, 220, 0.08)',
          'rgba(132, 34, 220, 0.03)',
          'rgba(132, 34, 220, 0.00)',
        ]}
        start={{ 
          x: 0, y: 0 
        }}
        end={{ 
          x: 1, y: 1 
        }}
        style={StyleSheet.absoluteFillObject}
      />

      <Text style={styles.title}>
        What's your email or username?
      </Text>

      <Input
        iconsName="mail"
        styleView={styles.input}
        labelColor="black"
        placeholder="hello@example.com"
      />

      <View style={styles.card}>
        <Text style={styles.cardTitle}>No spam guaranteed</Text>
        <Text style={styles.cardText}>
          We hate cluttered inboxes as much as you do. Only the good stuff.
        </Text>
      </View>

      <CustomButton
        style={styles.button}
        onPress={() => {}}
        title="Continue"
        iconsName="arrow-right"
      />
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    padding: 30,
    paddingTop: 50,
    backgroundColor: 'white',
  },
  title: {
    fontSize: 40,
    fontWeight: 'bold',
    marginBottom: 80,
    color: 'black',
  },
  input: {
    width: '100%',
    marginTop: 20,
  },
  card: {
    marginTop: 20,
    backgroundColor: '#F1F1F1',
    paddingHorizontal: 20,
    paddingVertical: 20,
    borderRadius: 20,
  },
  cardTitle: {
    fontWeight: '700',
    color: 'black',
  },
  cardText: {
    fontSize: 14,
    marginTop: 10,
    color: '#1A1C1C',
  },
  button: {
    marginTop: 'auto',
    backgroundColor: '#8422DC',
  },
});