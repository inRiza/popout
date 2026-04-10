import Input from '@/components/ui/input';
import { StyleSheet, View, Text } from 'react-native';
import CustomButton from '@/components/ui/button';
import { LinearGradient } from 'expo-linear-gradient';

export default function RegisterScreen() {
  return (
    <View style={styles.container}>
      <LinearGradient
        colors={[
          'rgba(132, 34, 220, 0.08)',
          'rgba(132, 34, 220a, 0.03)',
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

      <View style={{
        flexDirection: 'column',
        justifyContent: 'center',
        height: '100%',
      }}>
        <Text style={[styles.title,
          {
            marginBottom: 20,
          }
        ]}>
          Create your new Pop!
        </Text>

        <Text style={{
            marginBottom: 30,
            fontSize: 30,
            color: 'black'
          }}>
          Enter your email to get started
        </Text>

        <View style={{
          flexDirection: 'column',
          gap: 4,
        }}>
          <Input
            styleView={styles.input}
            labelColor="black"
            placeholder="hello@example.com"
          />
          {/* <Input
            styleView={styles.input}
            labelColor="black"
            placeholder="Enter your password"
          /> */}
        </View>
        <View style={{
          marginTop: 120,
          gap: 10,
        }}>
          <CustomButton
            style={styles.button}
            onPress={() => {}}
            title="Continue"
            iconsName="arrow-right"
          />
          <CustomButton
            style={styles.buttonSecondary}
            onPress={() => {}}
            title="Continue with Google"
            colorTitle={'black'}
            // iconsName="arrow-right"
          />
        </View>
      </View>
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
    fontSize: 50,
    fontWeight: 'bold',
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
    backgroundColor: '#8422DC',
  },
  buttonSecondary: {
    backgroundColor: '#E2E2E2',
  },
});